#shareIt
We all know the problem. You need to share a file with a customer but it's to big to attach it on e mail. Or a customer from you have to send you confidential data and you don't trust cloud hosters.

Now, there is a solution: shareIt


```
███████╗██╗  ██╗ █████╗ ██████╗ ███████╗██╗████████╗
██╔════╝██║  ██║██╔══██╗██╔══██╗██╔════╝██║╚══██╔══╝
███████╗███████║███████║██████╔╝█████╗  ██║   ██║   
╚════██║██╔══██║██╔══██║██╔══██╗██╔══╝  ██║   ██║   
███████║██║  ██║██║  ██║██║  ██║███████╗██║   ██║   
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝   ╚═╝   
                                                    
```


##Requirements
* Docker


##Usage
```
docker run -e ADMIN_PW=yourSecretPW -p 9000:9000 n3r0ch/shareit
```

Create a new room under `http://localhost:9000/admin`. Afterward you can simply upload files. Every user who are knowing the room URL (`http://localhost:9000/room#[secret-id]`) are permitted to access, edit and delete all files in the given room.


##API
It's also possible to access "shareIt" over a REST-API.

| Method | Route | Description |
|---|---|---|
| GET | /api/room/{id} | Returns the current content of a room |
| POST | /api/room | Creates a new room (no parameter needed) |
| GET | /api/room/{roomid}/file/{filename} | Downloads the given file |
| DELETE | /api/room/{roomid}/file/{filename} | Deletes the given file |
| POST | /api/room/{roomid}/file | Upload a file. If the given filename already exists, the old one will be overwritten |


##Volumes
* `/data`

All your uploaded files are stored under `/data`


##Exposed Ports
* 9000


##Bugs and Issues
If you have any problems with this image, feel free to open a new issue in our issue tracker [https://github.com/n3r0-ch/shareIt/issues](https://github.com/n3r0-ch/shareIt/issues)


##License
This image is licensed under the MIT License. The full license text is available under [https://github.com/n3r0-ch/shareIt/blob/master/LICENSE](https://github.com/n3r0-ch/shareIt/blob/master/LICENSE).