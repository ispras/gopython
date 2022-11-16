class TestClassInt:
	def __init__(self, arg):
		if not isinstance(arg, int):
			raise ValueError('arg should be int')

class TestClassFloat:
	def __init__(self, arg):
		if not isinstance(arg, float):
			raise ValueError('arg should be float')

class TestClassPyobj:
	def __init__(self, arg):
		if not isinstance(arg, TestClassInt):
			raise ValueError('arg should be TestClassInt')
