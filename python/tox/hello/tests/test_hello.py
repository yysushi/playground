import unittest

import hello


class TestHello(unittest.TestCase):

    def test_hello(self):
        self.assertEqual(hello.LANG, 'JP')
