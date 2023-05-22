from setuptools import setup

setup(
    name='backend',
    packages=['main','db','gpio'],
    include_package_data=True,
    install_requires=[
        'flask',
    ],
)