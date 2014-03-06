/**
 * Twist v0.1-dev
 *
 * (c) Harry Lawrence
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

/**
 * @package php
 * @author Harry Lawrence
 *
 * Twist's PHP object literal for using PHP
 * code inside of the editor
 */
var php = {
	namespace : "",
	class 	: "",

    /**
     * Creates a new class with optional
     * namespace, replaces the namespace
     * param to class if it's not set
     *
     * @param String namespace
     * @param String class name
     *
     * @return String class snippet
     */
	newClass : function(namespace, className) {
        if (arguments.length == 1) {
            this.class = namespace;

            snippet = "<?php\n\nclass " + namespace + "\n{\n\n}";
        } else {
            this.class     = className
            this.namespace = namespace;

            snippet = "<?php\n\nnamespace " + namespace + ";\n\nclass " + className + "\n{\n\n}";
        }

		return snippet;
	},

    /**
     * Creates a new function
     *
     * @param String function name
     *
     * @return String function snippet
     */
	newFunction : function(name) {
		return "function " + name + " ()\n{\n\n}";
	},

    /**
     * Alias for 'newFunction'
     *
     * @param String function name
     *
     * @return String function snippet
     */
	newFunc : function(name) {
		return this.newFunction(name)
	}
}
