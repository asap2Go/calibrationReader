# calibrationReader
 reads characteristics information from a2l and fills it with the data from a hex file.
 At least that is the plan. 
 Currently it can parse a2l-files as well as the corresponding IntelHex32 or Motorola S19 files. And it is quite fast at that (< 1s).
 
 But it still lacks the last bit of work which is implementing the methods for axis_pts, axis_descr, record_layout and fnc_values
 in order to understand the memory layout and position of a specific characteristic.
 This is somewhat of a convoluted mess in the a2l standard due to its historic growth and will be implemented when I have a little more spare time.
 
 I am a mechanical engineer, so any code you see might not be up to the standards of modern software development ;)
 Feedback is always much appreciated.
