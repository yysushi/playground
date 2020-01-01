import unittest
from unittest import mock

import requests

import get_ip


@mock.patch('get_ip.requests.get')
class TestGetIP(unittest.TestCase):

    def test_get_ip(self, get_patch):
        # given
        ip_mock = mock.MagicMock()
        ip_mock.json.return_value = {'origin': '8.8.8.8'}
        get_patch.return_value = ip_mock
        # when
        ip = get_ip.get_ip()
        # then
        self.assertEqual('8.8.8.8', ip)
        get_patch.assert_called_once_with("http://httpbin.org/ip")

    def test_get_ip_failure(self, get_patch):
        # given
        get_patch.side_effect = requests.exceptions.ConnectTimeout
        # when, then
        with self.assertRaises(get_ip.GetIPError):
            get_ip.get_ip()
        get_patch.assert_called_once_with("http://httpbin.org/ip")
