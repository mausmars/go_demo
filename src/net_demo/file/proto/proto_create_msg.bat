E:
set Project_Path=E:\vm_share\demo\net_demo
set Proto_Tool_Path=%Project_Path%\file\tool\proto_3.2.0\windows
set Proto_File_Path=%Project_Path%\file\proto

set Target_Path_Example1=%Project_Path%\dynamic\msg

cd  %Proto_File_Path%
%Proto_Tool_Path%/protoc.exe -I=%Proto_File_Path% --go_out=%Target_Path_Example1% %Proto_File_Path%/login.proto
