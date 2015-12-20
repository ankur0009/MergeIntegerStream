# MergeIntegerStream
Merges two integer streams and returns the lowest number from the merged list. For each invocation, the input streams are invoked to get two integers, and the lowest number from the merged cumulative list is returned.

Usage Notes
http://localhost:<port>/quiz/merge?stream1=<stream_name_1>&stream2=<stream_name_2>

Other Notes
- For a single instance of the application, if you use different pairs of streams, the data is merged for all the streams. In other words, we don't provide the lowest number per stream pair.
