/*

Package logdur provides a wrapper for log.Print to measure
the duration of a function's execution. It can be used
conveniently in conjunction with "defer". The statement

  defer logdur.Start().Print()

prints the time elapsed between the call to Start (when the
defer statement executes) and the call to Print (when the
function returns).

Output is formatted with package text/template.

*/
package logdur
