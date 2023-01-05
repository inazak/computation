use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $d = -1;
my $e = -1;
my $o = -1;

$d = LO();
$e = HI();
$o = DLatch($d, $e);

printf("//DLatch($f, $f) -> $f \n", $d, $e, $o);
Write();

$d = HI();
$e = HI();
$o = DLatch($d, $e);

printf("//DLatch($f, $f) -> $f \n", $d, $e, $o);
Write();


