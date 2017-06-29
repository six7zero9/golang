# Best Practices
https://www.youtube.com/watch?v=8D3Vmm1BGoY

### Twelve Best Practices

1. Avoid nesting by handling errors first
2. Avoid repition when possible
3. Important code goes first
4. Document your code (doc.go in convention)
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent
10. Avoid using concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks


### Other notes

1. early return on error
2. d.r.y (derp)
3. make a license if open sourced
4. document code (comment)
5. put helper methods at the end so code is easier to read
6. use http://godoc.org/ for help with available funcs/methods and viewing of source code for all go packages
7. doc.go is convention for documentation
8. use interfaces to avoid coupled code
9. avoid over using concurrency. especially for API's
