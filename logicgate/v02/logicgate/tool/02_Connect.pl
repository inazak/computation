use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

Connect(HI(), 1000);
Connect(HI(), 1001);
Connect(LO(), 1002);
Connect(LO(), 1003);
Connect(HI(), 1004);
Connect(HI(), 1005);
Connect(LO(), 1006);
Connect(LO(), 1007);
Write();

