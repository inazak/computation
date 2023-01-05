use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $a    = -1;
my $b    = -1;
my $cin  = -1;
my $s    = -1;
my $cout = -1;

$a   = LO();
$b   = LO();
$cin = LO();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = HI();
$b   = LO();
$cin = LO();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = LO();
$b   = HI();
$cin = LO();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = HI();
$b   = HI();
$cin = LO();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = LO();
$b   = LO();
$cin = HI();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = HI();
$b   = LO();
$cin = HI();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = LO();
$b   = HI();
$cin = HI();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

$a   = HI();
$b   = HI();
$cin = HI();
($s, $cout) = Fadd($a, $b, $cin);
printf("//Fadd($f, $f, $f) -> $f, $f \n", $a, $b, $cin, $s, $cout);
Write();

