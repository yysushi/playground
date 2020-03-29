import ipaddress
import unittest
import unittest.mock
import urllib.error

import mypkg


class TestPerson(unittest.TestCase):

    def setUp(self):
        self.person = mypkg.Person('john')

    # def tearDown(self):
    #     print(a.Person.SOURCE)

    def test_greet(self):
        actual = self.person.greet()
        expected = 'hello john'
        self.assertEqual(actual, expected)

    @unittest.mock.patch('time.sleep', autospec=True)
    def test_sleep(self, sleep_mock):
        self.person.sleep(12.2)
        sleep_mock.assert_called_once_with(12.2)

    def test_is_where(self):
        addresses = [
            self.person.is_where_born(),
            self.person.is_where_living(),
            self.person.is_where_living2(),
        ]
        for address in addresses:
            try:
                ipaddress.IPv4Address(address)
            except ValueError:
                msg = "Invalid address like %s".format(address)
                self.fail(msg)

    # @unittest.mock.patch('mypkg.external.SOURCE', new='https://nowhere')
    @unittest.mock.patch('mypkg.person.SOURCE', new='https://nowhere')
    def test_is_where_failure(self):
        # exception
        # new_person = mypkg.Person('mary')
        with self.assertRaises(urllib.error.URLError):
            print(self.person.is_where_born())

    @unittest.mock.patch('mypkg.person.get_active_source', autospec=True)
    def test_is_where_failure2(self, mock_active_source):
        mock_active_source.return_value = 'https://nowhere'
        # exception
        with self.assertRaises(urllib.error.URLError):
            print(self.person.is_where_living())
        mock_active_source.assert_called_once()

    @unittest.mock.patch.object(mypkg.person.Discovery, 'active_source', autospec=True)
    def test_is_where_failure3(self, mock_active_source):
        mock_active_source.return_value = 'https://nowhere'
        # exception
        with self.assertRaises(urllib.error.URLError):
            print(self.person.is_where_living2())
        mock_active_source.assert_called_once()


if __name__ == '__main__':
    unittest.main()
