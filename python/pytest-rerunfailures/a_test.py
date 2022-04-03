import pytest


class TestA:

    @classmethod
    def setup_class(cls):
        print('AAA')

    def setup(self):
        print('AA')

    @pytest.mark.flaky(reruns=5, reruns_delay=2)
    def test_a(self):
        raise Exception('failed')
