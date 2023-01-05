use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $i  = -1;
my $q1 = -1;
my $q2 = -1;

$i = LO();
($q1, $q2) = BDec($i);
printf("//BDec($f) -> $f, $f \n", $i, $q1, $q2);
Write();

$i = HI();
($q1, $q2) = BDec($i);
printf("//BDec($f) -> $f, $f \n", $i, $q1, $q2);
Write();

