import unittest

from poetry_demo.request import google


class TestG(unittest.TestCase):

    def test_google(self):
        self.assertEqual(301, google())
