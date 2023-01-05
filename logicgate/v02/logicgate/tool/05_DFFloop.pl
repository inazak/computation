use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $d = 1000;
my $clock = CLOCK();

my $o = DFF($d, $clock);
my $p = Not($o);
Connect($p, $d);

printf("//DFF($f, $f) -> $f : NOT($f) -> $f : $f -> $f\n", $d, $clock, $o, $o, $p, $p, $d);
Write();



