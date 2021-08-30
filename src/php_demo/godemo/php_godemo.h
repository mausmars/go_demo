/* godemo extension for PHP */

#ifndef PHP_GODEMO_H
# define PHP_GODEMO_H

extern zend_module_entry godemo_module_entry;
# define phpext_godemo_ptr &godemo_module_entry

# define PHP_GODEMO_VERSION "0.1.0"

# if defined(ZTS) && defined(COMPILE_DL_GODEMO)
ZEND_TSRMLS_CACHE_EXTERN()
# endif

#endif	/* PHP_GODEMO_H */

PHP_FUNCTION(godemoHello);
PHP_FUNCTION(godemoAdd);