# Singolang
Singolang is a library to interact with Singularity containers in Go. It is modeled from Spython. Designed for use with Singularity 3+

# Instance
Instance works differently from Spython, you need to create an Instance using instance.GetInstance() before calling Start or Stop. You cannot Start or Stop instances that were not created by this method.