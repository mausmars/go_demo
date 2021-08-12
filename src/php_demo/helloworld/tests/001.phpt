--TEST--
Check if helloworld is loaded
--SKIPIF--
<?php
if (!extension_loaded('helloworld')) {
	echo 'skip';
}
?>
--FILE--
<?php
echo 'The extension "helloworld" is available';
?>
--EXPECT--
The extension "helloworld" is available
