<h3 align="center">Description</h3>

The repository holds implementation of parallel and sequential inverted index. To start working with solution follow simple instructions below:

<b><i>1.Clone the repository</i></b>

`git clone git@github.com:Illia-tsar/course_work_parallel_computing.git`

<b><i>2.Change the directory</i></b>

`cd course_work_parallel_computing/`

<b><i>3.Create new directory</i></b>

`mkdir data`

Next, add documents to this directory, so that inverted index has data to work on.

<b><i>4.Run parallel inverted index with 2 goroutines and show execution time</i></b>

`go run . -n 2 -pb=true -t=true`

<h5 align="center">Parameters:</h5>
* -t - show execution time(defaults to true)
* -pb - parallel build(defaults to true)
* -n - the number of goroutines(defaults to 1)
* -s - search term in index, or not(defaults to false)
* -sw - the word to search(defaults to "")