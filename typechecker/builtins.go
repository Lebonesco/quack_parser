package typechecker

var BUIILT_IN_CLASSES = 
` class Obj() {
    def PRINT() { }
    def STR() { }
    def EQUALS(other: Obj): Boolean {} // Default is object identity
 }

 class Nothing() { }

 class String() {
    def PLUS(other: String): String { }      // +
    def EQUALS(other: String): Boolean { }   // ==
    def ATMOST(other: String): Boolean { }   // <=
    def LESS(other: String): Boolean { }     // <
    def ATLEAST(other: String): Boolean { }  // >=
    def MORE(other: String): Boolean { }     // >
 }

 class Boolean() { }

 class Int() {
    def PLUS(right: Int): Int {}   // this + right
    def TIMES(right: Int): Int {}  // this * right
    def MINUS(right: Int): Int {}  // this - right
    def DIVIDE(right: Int): Int {}    // this / right
    def ATMOST(other: Int): Boolean { }     // <=
    def LESS(other: Int): Boolean { }       // <
    def ATLEAST(other: Int): Boolean { }    // >=
    def MORE(other: Int): Boolean { }       // >
 }
`