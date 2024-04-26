# Create .zip
# Create sample text files


echo "This is file1.txt" > file1.txt
echo "This is file2.txt" > file2.txt
echo "This is file3.txt" > file3.txt
echo "this is WindowsActiveDirectoryDatabaseFile" > WindowsActiveDirectoryDatabaseFile.txt

# Create a directory and a file inside it
mkdir fileDirectory
echo "this is WindowsUserCustomDestinationsJumpLists" > fileDirectory/WindowsUserCustomDestinationsJumpLists.txt

echo "this is file4.txt" > fileDirectory/file4.txt

mkdir fileDirectory/nestedDirectory
echo "this is WindowsUserCustomDestinationsJumpLists inside a directory" > fileDirectory/nestedDirectory/WindowsUserCustomDestinationsJumpLists.txt
echo "this is file5.txt" > fileDirectory/nestedDirectory/file5.txt

# Create a ZIP archive containing the text files
zip test.zip *.txt fileDirectory/*.txt fileDirectory/nestedDirectory/*
