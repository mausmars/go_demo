--TEST--
helloworld_test1() Basic test
--SKIPIF--
<?php
if (!extension_loaded('helloworld')) {
	echo 'skip';
}
?>
--FILE--
<?php
$ret = helloworld_test1();

var_dump($ret);
?>
--EXPECT--
The extension helloworld is loaded and working!
NULL
