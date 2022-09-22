# Modules

This chapter will explain how walnut store modules.

## Kind of Modules

* local
* linked
* mirrored

### local-Modules

Local modules are stored inside the volume of the walnut-store service.
If walnut runs inside docker using docker-compose this volume is mapped onto the local filesystem.
Inside kubernetes walnut can store the local modules in several different ways. [kubernetes volume](https://kubernetes.io/docs/concepts/storage/volumes/)


### linked-Modules

Linked modules are not stored inside walnut.
They are only referenced by a url inside the walnutdb.
If a build request them the store service tries to load the modules from the url.
An error is send if the module can not be loaded, the errors are explained inside the errors chapter of the walnut docs.
Linked Modules can ether be the url to a http-based index-registry or a publicly available fileserver.

### mirrored-Modules

Mirrored modules are modules that are loaded initialy when they are requested, later if they are requested walnut checks the original url for a newer version of the module and if it finds them they are getting downloaded to walnut as well.


