# MergeIntegerStream
Merges two integer streams and returns the lowest number from the merged list. For each invocation, the input streams are invoked to get two integers, and the lowest number from the merged cumulative list is returned.

Usage Notes
- http://localhost:<port>/quiz/merge?stream1=<stream_name_1>&stream2=<stream_name_2>

Time complexity Notes
- The values from the input streams are merged into a single list. For every invocation the list is searched for the lowest value and the value is then removed from the list. The time complexity is O(1) for insertion, and O(n) for search.

- However, this performance can be improved by using an array based Min Heap to store the values. The time complexity with a Min Heap would be O(log n) for insertion and O(1) for search. THIS IS NOT DONE YET.
 
Memory Usage Notes
- When a value from the merged list is generated, it's spot is marked as empty and reused for the next input value.

Other Notes
- Duplicate values from the streams are NOT ingored.

- For a single instance of the application, if you use different pairs of streams, the data is merged for all the streams. In other words, we don't provide the lowest number per stream pair.
