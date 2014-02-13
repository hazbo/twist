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

var php = {
	phpNamespace : "",
	phpClassName : "",

	newClass : function(namespace, className) {
		this.phpClassName = className;
		this.phpNamespace = namespace;
		return "<?php\n\nnamespace " + namespace + ";\n\nclass " + className + "\n{\n\n}";
	},

	newFunction : function(name) {
		return "function " + name + " ()\n{\n\n}";
	},

	newFunc : function(name) {
		return this.newFunction(name)
	}
}