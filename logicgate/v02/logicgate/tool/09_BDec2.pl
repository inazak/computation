use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $s1 = -1;
my $s2 = -1;
my $q1 = -1;
my $q2 = -1;
my $q3 = -1;
my $q4 = -1;

$s1 = LO();
$s2 = LO();
($q1, $q2, $q3, $q4) = BDec2($s1, $s2);
printf("//BDec2($f,$f) -> $f, $f, $f, $f \n", $s1, $s2, $q1, $q2, $q3, $q4);
Write();

$s1 = HI();
$s2 = LO();
($q1, $q2, $q3, $q4) = BDec2($s1, $s2);
printf("//BDec2($f,$f) -> $f, $f, $f, $f \n", $s1, $s2, $q1, $q2, $q3, $q4);
Write();

$s1 = LO();
$s2 = HI();
($q1, $q2, $q3, $q4) = BDec2($s1, $s2);
printf("//BDec2($f,$f) -> $f, $f, $f, $f \n", $s1, $s2, $q1, $q2, $q3, $q4);
Write();

$s1 = HI();
$s2 = HI();
($q1, $q2, $q3, $q4) = BDec2($s1, $s2);
printf("//BDec2($f,$f) -> $f, $f, $f, $f \n", $s1, $s2, $q1, $q2, $q3, $q4);
Write();

