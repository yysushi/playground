import ipaddress
import unittest
import unittest.mock
import urllib.error

import a


# NOTE: many return_value... I should learn better one.
def get_url_opened_mock(value=None, exc=None):
    if value is not None:
        fd_mock = unittest.mock.MagicMock(**{
            'read.return_value': bytes('{"origin": "%s"}' % value, 'utf-8'),
        })
    elif exc is not None:
        fd_mock = unittest.mock.MagicMock(**{
            'read.side_effect': exc,
        })
    return unittest.mock.MagicMock(**{
        '__enter__.return_value': fd_mock,
    })


class TestPerson(unittest.TestCase):

    def setUp(self):
        self.person = a.Person('john')

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

    def test_is_where_failure(self):
        a.Person.SOURCE = 'nowhere'
        with self.assertRaises(urllib.error.URLError):
            self.person.is_where()

    @unittest.mock.patch('urllib.request.urlopen')
    def test_is_where_with_mock(self, urlopen_mock):
        urlopen_mock.return_value = \
            get_url_opened_mock(value='8.8.8.8, 1.1.1.1')
        address = self.person.is_where()
        self.assertEqual(address, '8.8.8.8')
        urlopen_mock.call_once_with('http://httpbin.org/ip')

    @unittest.mock.patch('urllib.request.urlopen')
    def test_is_where_failure_with_mock(self, urlopen_mock):
        urlopen_mock.return_value = \
            get_url_opened_mock(exc=urllib.error.URLError('no such endpoint'))
        with self.assertRaises(urllib.error.URLError):
            self.person.is_where()
        urlopen_mock.call_once_with('http://httpbin.org/ip')


if __name__ == '__main__':
    unittest.main()
