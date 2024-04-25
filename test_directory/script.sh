# Create .zip
# Create sample text files


echo "This is file1.txt" > file1.txt
echo "This is file2.txt" > file2.txt
echo "This is file3.txt" > file3.txt
echo "WindowsActiveDirectoryDatabaseFile" > WindowsActiveDirectoryDatabaseFile.txt

# Create a directory and a file inside it
mkdir file
echo "WindowsUserCustomDestinationsJumpLists" > file/WindowsUserCustomDestinationsJumpLists.txt

# Create a ZIP archive containing the text files
zip test.zip *.txt file/*.txt
