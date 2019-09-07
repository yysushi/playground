import ipaddress
import json
import unittest
import unittest.mock
import urllib.error

import a


class TestPerson(unittest.TestCase):

    def setUp(self):
        self.person = a.Person('john')

    def tearDown(self):
        print(a.Person.SOURCE)

    def test_greet(self):
        actual = self.person.greet()
        expected = 'hello john'
        self.assertEqual(actual, expected)

    @unittest.mock.patch('time.sleep')
    def test_sleep(self, sleep_mock):
        self.person.sleep(2.2)
        sleep_mock.call_once_with(2.2)

    def test_is_where(self):
        address = self.person.is_where()
        try:
            ipaddress.IPv4Address(address)
        except ValueError:
            msg = "Invalid address like %s".format(address)
            self.fail(msg)

    # @unittest.mock.patch('a.Person', wraps=a.Person, autospec=True)
    @unittest.mock.patch('a.Person', wraps=a.Person)
    def test_is_where_failure(self, person_mock):
        print("")
        print("hoge", a.Person.SOURCE, dir(a.Person))
        person_mock.SOURCE = 'nowhere'
        print("fuga", a.Person.SOURCE)
        with self.assertRaises(urllib.error.URLError):
            self.person.is_where()

    # @unittest.mock.patch('urllib.request.urlopen')
    # def test_is_where_with_mock(self, urlopen_mock):
    #     unittest.mock.patch.object(urlopen_mock, '')
    #     data = bytes(json.dumps({'origin': '8.8.8.8, 1.1.1.1'}), 'utf-8')
    #     unittest.mock.patch.object(urlopen_mock, 'read', return_value=data)
    #     address = self.person.is_where()
    #     self.assertEqual(address, '8.8.8.8')
    #     urlopen_mock.call_once_with('http://httpbin.org/ip')

    # @unittest.mock.patch('urllib.request.urlopen', create=True)
    # def test_is_where_with_mock(self, urlopen_mock):
    # def test_is_where_with_mock(self):
    #     data = bytes(json.dumps({'origin': '8.8.8.8, 1.1.1.1'}), 'utf-8')
    #     fd_mock = unittest.mock.MagicMock(**{
    #         'read.return_value': data,
    #     })
    #     unittest.mock.patch.object(
    #         urllib.request.urlopen, '__enter__', return_value=fd_mock)
    #     # unittest.mock.patch.object(
    #     #     urllib.request.urlopen, '__enter__', return_value=fd_mock)
    #     address = self.person.is_where()
    #     self.assertEqual(address, '8.8.8.8')
    #     urlopen_mock.call_once_with('http://httpbin.org/ip')

    # @unittest.mock.patch('urllib.request.urlopen')
    # def test_is_where_failure_with_mock(self, urlopen_mock):
    #     urlopen_mock.return_value = \
    #         get_url_opened_mock(exc=urllib.error.URLError('no such endpoint'))
    #     with self.assertRaises(urllib.error.URLError):
    #         self.person.is_where()
    #     urlopen_mock.call_once_with('http://httpbin.org/ip')


if __name__ == '__main__':
    unittest.main()
