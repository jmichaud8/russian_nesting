# russian_nesting

**Description** </br>
In this coding challenge, you are tasked with solving the "Russian Doll Envelopes" problem. Given a 2D array of integers, envelopes, where envelopes[i] = [wi, hi] represents the width and the height of an envelope, your goal is to determine the maximum number of envelopes you can nest inside each other. An envelope can fit into another only if both its width and height are strictly greater than those of the envelope it fits into. Note that you cannot rotate the envelopes.

**Requirements** </br>
*Programming Language*: Solutions must be implemented in Go. This will test the participants' proficiency with Go-specific paradigms and data structures. </br>
*Code Style*: Adherence to Go's standard coding style (as outlined in Go's Effective Go document) is expected. Readability and proper formatting are part of the evaluation criteria. </br>
*Documentation*: Inline comments and a brief documentation on how the solution works and why certain approaches were chosen can provide insight into the thought process of the coder, which is beneficial during scoring. </br>

**Constraints** </br>
*Standard Input*: The algorithm should read from standard input and not expect input from files or other sources. </br>
*Input Format*: Input will be provided on the command line exactly in the format [[w1,h1],[w2,h2],...,[wn,hn]], where each pair [wi, hi] corresponds to the width and height of an envelope. </br>
*Parsing Requirement*: Participants must write their code to parse this input format directly from standard input (stdin). </br>
*Envelope Count*: The number of envelopes, n, derived from the input, must comply with the range 1 <= n <= 10^5. </br>
*Dimension Constraints*: Each envelope's width and height, wi and hi, will satisfy 1 <= wi, hi <= 10^5. </br>
*No Partial or Malformed Data*: It is assumed that all input data will be well-formed and complete. </br>

**Output**
* A single integer should be returned to the console on how many envelopes could be nested using the provided input  </br>

**Scoring**</br>
*Correctness and Completeness (50 points)* </br>
* Full marks are awarded if the solution passes all provided test cases, including edge cases. </br>
* Partial marks are given if the solution passes only some test cases. </br>
* No points are awarded if the solution fails to compile or crashes on the majority of test cases. </br>

*Performance and Efficiency (30 points)* </br>
* Full marks are awarded if the solution’s runtime is the fastest out of all submissions. </br>
* Partial marks are given if the solution is not the slowest. </br>
* No points are awarded if the solution exhibits poor performance - is the slowest solution. </br>

*Code Quality and Style (20 points)* </br>
* Full marks are awarded for well-structured code with clear, meaningful comments, proper variable naming, and adherence to Go's formatting and style guidelines. </br>
* Partial marks are given if the code is mostly well-written but lacks comments or has minor style issues. </br>
* Few or no points are awarded for code that is difficult to read, poorly structured, or lacking in documentation. </br>

 </br>
  </br>

Examples </br>
_Example 1_ </br>
**Input**: envelopes = [[5,4],[6,4],[6,7],[2,3]] </br>
**Output**: 3 </br>
**Explanation**: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]). </br>


_Example 2_ </br>
*Input*: envelopes = [[1,1],[1,1],[1,1]] </br>
*Output*: 1 </br>
*Explanation*: Only one envelope can be placed as all are of the same size. </br>

_Example 3_ </br>
*Input*: [[10,8],[1,12],[6,15],[2,18]] </br>
*Output*: 2 </br>
Explanation: The largest number of envelopes you can nest is 2 ([1,12] => [6,15] or [1,12] => [2,18]). </br>
