use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $a0   = -1;
my $a1   = -1;
my $a2   = -1;
my $a3   = -1;
my $b0   = -1;
my $b1   = -1;
my $b2   = -1;
my $b3   = -1;
my $cin  = -1;
my $s0   = -1;
my $s1   = -1;
my $s2   = -1;
my $s3   = -1;
my $cout = -1;

$a0  = LO();
$a1  = LO();
$a2  = LO();
$a3  = LO();
$b0  = LO();
$b1  = LO();
$b2  = LO();
$b3  = LO();
$cin = LO();
($s0,$s1,$s2,$s3, $cout) = Fadd4($a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin);
printf("//Fadd4($f,$f,$f,$f, $f,$f,$f,$f, $f) -> $f,$f,$f,$f, $f\n", $a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin, $s0,$s1,$s2,$s3, $cout);
Write();

$a0  = HI();
$a1  = LO();
$a2  = HI();
$a3  = LO();
$b0  = HI();
$b1  = LO();
$b2  = HI();
$b3  = LO();
$cin = LO();
($s0,$s1,$s2,$s3, $cout) = Fadd4($a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin);
printf("//Fadd4($f,$f,$f,$f, $f,$f,$f,$f, $f) -> $f,$f,$f,$f, $f\n", $a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin, $s0,$s1,$s2,$s3, $cout);
Write();

$a0  = LO();
$a1  = HI();
$a2  = HI();
$a3  = LO();
$b0  = HI();
$b1  = LO();
$b2  = LO();
$b3  = HI();
$cin = HI();
($s0,$s1,$s2,$s3, $cout) = Fadd4($a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin);
printf("//Fadd4($f,$f,$f,$f, $f,$f,$f,$f, $f) -> $f,$f,$f,$f, $f\n", $a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin, $s0,$s1,$s2,$s3, $cout);
Write();

$a0  = HI();
$a1  = HI();
$a2  = HI();
$a3  = HI();
$b0  = HI();
$b1  = HI();
$b2  = HI();
$b3  = HI();
$cin = HI();
($s0,$s1,$s2,$s3, $cout) = Fadd4($a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin);
printf("//Fadd4($f,$f,$f,$f, $f,$f,$f,$f, $f) -> $f,$f,$f,$f, $f\n", $a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $cin, $s0,$s1,$s2,$s3, $cout);
Write();

