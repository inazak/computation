use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $d0    = -1;
my $d1    = -1;
my $clock = -1;
my $load  = -1;
my $clear = -1;
my $q0    = -1;
my $q1    = -1;

$d0    = LO();
$d1    = LO();
$clock = CLOCK();
$load  = LO();
$clear = LO();
($q0, $q1) = Counter2($d0, $d1, $clock, $load, $clear);
printf("//Counter2($f, $f, $f, $f, $f) -> $f, $f \n", $d0, $d1, $clock, $load, $clear, $q0, $q1);
Write();

$d0    = LO();
$d1    = LO();
$clock = CLOCK();
$load  = LO();
$clear = HI();
($q0, $q1) = Counter2($d0, $d1, $clock, $load, $clear);
printf("//Counter2($f, $f, $f, $f, $f) -> $f, $f \n", $d0, $d1, $clock, $load, $clear, $q0, $q1);
Write();

$d0    = HI();
$d1    = HI();
$clock = CLOCK();
$load  = HI();
$clear = LO();
($q0, $q1) = Counter2($d0, $d1, $clock, $load, $clear);
printf("//Counter2($f, $f, $f, $f, $f) -> $f, $f \n", $d0, $d1, $clock, $load, $clear, $q0, $q1);
Write();

