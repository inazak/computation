use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $a = -1;
my $b = -1;
my $s = -1;
my $o = -1;

$a = LO();
$b = LO();
$s = LO();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = LO();
$b = HI();
$s = LO();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = HI();
$b = LO();
$s = LO();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = HI();
$b = HI();
$s = LO();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();


$a = LO();
$b = LO();
$s = HI();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = LO();
$b = HI();
$s = HI();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = HI();
$b = LO();
$s = HI();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

$a = HI();
$b = HI();
$s = HI();
$o = Mux($a, $b, $s);
printf("//Mux($f, $f, $f) -> $f \n", $a, $b, $s, $o);
Write();

