# Plasticine

A tiny ray tracer made using Go following the
[Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)
book. One part of my adventures with computer graphics, and how to render them
in software.

This one does go beyond by running the ray tracing in parallel on all CPU cores.

## Running

```shell
$ mkdir out
$ go run ./...
```

An output image will be created at `out/out.png`.

![A render of 3 spheres](./render.png)

The render was done with the resolution of `1440x810` with 256 samples and 96
light bounces (`main.maxDepth`). Took 22.395 seconds on my Ryzen 7 5800U.

## Scene

The scene is currently hardcoded in [object/world.go](./object/world.go), and
only supports spheres as the meshes. Currently implemented shaders include
[diffuse](./object/lambertian.go), [metal](./object/metal.go) and
[dielectric](./object/dielectric.go).
