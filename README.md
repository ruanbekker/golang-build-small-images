# golang-build-small-images
Demo of Building Small Golang Images

## About
So we thinking lets go with alpine right? Yeah sure lets build a small, lekke, but useless app running on go with alpine.

Let's package our app to an image:

```
❯ docker build -t mygolangapp:using-alpine .
```

Inspect the size of our image, as you can see it being 310MB

```
❯ docker images "mygolangapp:*"
REPOSITORY          TAG                 IMAGE ID            CREATED              SIZE
mygolangapp         using-alpine        eea1d7bde218        About a minute ago   310MB
```

Just make sure it actually works:

```
❯ docker run mygolangapp:using-alpine
Number of words: 16
Selected index number: 11
Selected word is: mountain lion
```

But for something just returning random selected text, 310MB is a bit crazy.

## Multi Stage Builds

As Go binaries are self-contained, we can make use of docker's multi stage builds, where we can build our application on alpine and use the binary on a scratch image:

Build it:

```
❯ docker build -t mygolangapp:using-multistage -f Dockerfile.multi .
```

Notice that the image is only 2.01MB, say w000t!

```
❯ docker images "mygolangapp:*"
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
mygolangapp         using-multistage    31474c61ba5b        15 seconds ago      2.01MB
mygolangapp         using-alpine        eea1d7bde218        2 minutes ago       310MB
```

Run the app:

```
❯ docker run mygolangapp:using-multistage
Number of words: 16
Selected index number: 5
Selected word is: spider
```
