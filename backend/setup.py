from setuptools import setup

setup(
    name='backend',
    packages=['main','db'],
    include_package_data=True,
    install_requires=[
        'flask',
    ],
)