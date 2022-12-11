package minefield

/*
The minimum percentage of tiles that must be revealed by the initial patch that
is revealed during the minefield generation
*/
const initialPatchMinCoverage float32 = 0.1

/*
The maximum number of patches that will be checked, during the minefield
generation, before aborting trying to reveal an initial patch
*/
const initialPatchMaxIterations int = 10
