import * as THREE from "three";

class Object3D {
  name;
  position = [0, 0, 0];
  color;
  // Size is now a direct property of the object
  size;
  // Reference to the source object
  sourceObject;
  // Actual mesh for the object (no longer using points)
  mesh = null;

  constructor(
    name,
    sourceObject = null,
    pos = [0, 0, 0],
    color = 0xffffff,
    size = 1.0
  ) {
    this.name = name;
    this.sourceObject = sourceObject;
    this.position = pos;
    this.color = color;
    this.size = size;
  }

  // Update from source object if available
  updateFromSource() {
    if (
      this.sourceObject &&
      typeof this.sourceObject.getPosition === "function"
    ) {
      this.position = this.sourceObject.getPosition();
    }
    if (this.sourceObject && typeof this.sourceObject.getColor === "function") {
      this.color = this.sourceObject.getColor();
    }
    if (this.sourceObject && typeof this.sourceObject.getSize === "function") {
      this.size = this.sourceObject.getSize();
    }
    return this;
  }

  updatePosition(pos) {
    this.position = pos;
    return this;
  }

  updateColor(color) {
    this.color = color;
    return this;
  }

  updateSize(size) {
    this.size = size;
    return this;
  }
}

class SolarSystem {
  time;
  renderer;
  scene;
  camera;
  i = 0;
  objects = [];
  customAnimationCallback = null;
  minSize = 0.05;
  maxSize = 2.0;
  // New: scaling factor for physical size to visual size
  sizeScaleFactor = 1.0;
  // Use a higher detail level for larger planets
  sphereSegments = 32;

  constructor(time, objectsData = [], options = {}) {
    this.time = time;

    // Apply options
    if (options.minSize !== undefined) this.minSize = options.minSize;
    if (options.maxSize !== undefined) this.maxSize = options.maxSize;
    if (options.sizeScaleFactor !== undefined)
      this.sizeScaleFactor = options.sizeScaleFactor;
    if (options.sphereSegments !== undefined)
      this.sphereSegments = options.sphereSegments;

    // Create scene with space background
    this.scene = new THREE.Scene();
    this.scene.background = new THREE.Color(0x0a0a1a);

    // Setup renderer
    this.renderer = new THREE.WebGLRenderer({ antialias: true });
    this.renderer.setSize(options.x, options.y);
    document
      .querySelector(options.element)
      .appendChild(this.renderer.domElement);

    // Cartoon-style lighting
    const hemisphereLight = new THREE.HemisphereLight(0xffffbb, 0x080820, 1);
    this.scene.add(hemisphereLight);

    // Add visible sun with glowing effect
    this.createSun();

    // Camera setup
    this.camera = new THREE.PerspectiveCamera(60, 1, 0.1, 1000);
    this.camera.position.set(0, 15, 5);
    this.camera.lookAt(0, 0, 0);

    // Add starfield and objects
    this.addStarfield();
    if (objectsData.length > 0) this.addObjects(objectsData);

    window.addEventListener("resize", this.onWindowResize.bind(this), false);
  }

  createSun() {
    // Sun light source
    const sunColor = 0xffd700;
    const pointLight = new THREE.PointLight(sunColor, 100, 0, 2);
    pointLight.position.set(0, 0, 0);
    this.scene.add(pointLight);

    // Sun visual representation
    const sunSize = 0.25;
    const geometry = new THREE.SphereGeometry(sunSize, 32, 32);

    // Glowing core material
    const coreMaterial = new THREE.MeshPhongMaterial({
      color: sunColor,
      emissive: sunColor,
      emissiveIntensity: 1.5,
      specular: 0xffffff,
      shininess: 100
    });

    // Outer glow effect
    const glowMaterial = new THREE.MeshBasicMaterial({
      color: sunColor,
      transparent: true,
      opacity: 0.3,
      blending: THREE.AdditiveBlending
    });

    // Sun core mesh
    const sunCore = new THREE.Mesh(geometry, coreMaterial);

    // Glow effect mesh
    const sunGlow = new THREE.Mesh(
      new THREE.SphereGeometry(sunSize * 1.4, 32, 32),
      glowMaterial
    );

    // Outline mesh
    const outlineMesh = new THREE.Mesh(
      new THREE.SphereGeometry(sunSize * 1.1, 32, 32),
      new THREE.MeshBasicMaterial({
        color: 0x000000,
        side: THREE.BackSide
      })
    );

    // Combine all sun elements
    const sun = new THREE.Group();
    sun.add(sunCore);
    sun.add(sunGlow);
    sun.add(outlineMesh);
    this.scene.add(sun);

    // Animation properties
    sun.userData.phase = Math.random() * Math.PI * 2;
    sun.userData.baseSize = sunSize;
  }

  addStarfield() {
    const geometry = new THREE.BufferGeometry();
    const vertices = [];
    for (let i = 0; i < 1000; i++) {
      vertices.push(
        THREE.MathUtils.randFloatSpread(200),
        THREE.MathUtils.randFloatSpread(200),
        THREE.MathUtils.randFloatSpread(200)
      );
    }
    geometry.setAttribute(
      "position",
      new THREE.Float32BufferAttribute(vertices, 3)
    );
    const material = new THREE.PointsMaterial({
      color: 0xffffff,
      size: 0.05,
      transparent: true,
      opacity: 0.8
    });
    const stars = new THREE.Points(geometry, material);
    this.scene.add(stars);
  }

  // Method to add new objects to the system
  addObjects(objectsData) {
    // Add new objects to the array
    for (let data of objectsData) {
      this.addObject(
        data.name,
        data.sourceObject || null,
        data.pos || [0, 0, 0],
        data.color !== undefined ? data.color : 0xffffff,
        data.size !== undefined ? data.size : 1.0
      );
    }
  }

  // Method to add a single object with a source object reference
  addObject(
    name,
    sourceObject = null,
    pos = [0, 0, 0],
    color = 0xffffff,
    size = 1.0
  ) {
    const obj = new Object3D(name, sourceObject, pos, color, size);
    this.objects.push(obj);

    // Create a mesh for the object
    this.createObjectMesh(obj);

    return obj; // Return the added object
  }

  createObjectMesh(obj) {
    const visualSize = this.calculateVisualSize(obj.size);

    // Main geometry
    const geometry = new THREE.SphereGeometry(
      visualSize,
      this.sphereSegments,
      this.sphereSegments
    );

    // Cartoonish material with outline effect
    const mainMaterial = new THREE.MeshPhongMaterial({
      color: obj.color,
      specular: 0xffffff,
      shininess: 100,
      emissive: obj.color,
      emissiveIntensity: 0.2
    });

    const outlineMaterial = new THREE.MeshBasicMaterial({
      color: 0x000000,
      side: THREE.BackSide
    });

    const mainMesh = new THREE.Mesh(geometry, mainMaterial);
    const outlineMesh = new THREE.Mesh(
      new THREE.SphereGeometry(
        visualSize * 1.1,
        this.sphereSegments,
        this.sphereSegments
      ),
      outlineMaterial
    );

    mainMesh.add(outlineMesh);
    mainMesh.position.set(...obj.position);

    // Add subtle animation offset
    mainMesh.userData.phase = Math.random() * Math.PI * 2;

    this.scene.add(mainMesh);
    obj.mesh = mainMesh;
  }

  // Calculate visual size based on actual size
  calculateVisualSize(actualSize) {
    // Apply size scaling and clamp between min and max size
    return Math.max(
      this.minSize,
      Math.min(this.maxSize, actualSize * this.sizeScaleFactor)
    );
  }

  // Method to get an object by name
  getObject(name) {
    return this.objects.find(obj => obj.name === name);
  }

  // Method to remove an object by name
  removeObject(name) {
    const index = this.objects.findIndex(obj => obj.name === name);
    if (index !== -1) {
      const obj = this.objects[index];

      // Remove mesh from scene
      if (obj.mesh) {
        this.scene.remove(obj.mesh);
        obj.mesh.geometry.dispose();
        obj.mesh.material.dispose();
      }

      // Remove from objects array
      this.objects.splice(index, 1);
      return true;
    }
    return false;
  }

  // Set a custom animation callback
  setAnimationCallback(callback) {
    if (typeof callback === "function") {
      this.customAnimationCallback = callback;
    }
  }

  // Configure size range and scaling
  setSizeParameters(minSize, maxSize, scaleFactor) {
    if (minSize !== undefined) this.minSize = minSize;
    if (maxSize !== undefined) this.maxSize = maxSize;
    if (scaleFactor !== undefined) this.sizeScaleFactor = scaleFactor;

    // Update all object meshes with new size parameters
    this.updateAllMeshes();
  }

  // Update all meshes (e.g., after changing size parameters)
  updateAllMeshes() {
    for (const obj of this.objects) {
      // Remove old mesh
      if (obj.mesh) {
        this.scene.remove(obj.mesh);
        obj.mesh.geometry.dispose();
        obj.mesh.material.dispose();
      }

      // Create new mesh with updated parameters
      this.createObjectMesh(obj);
    }
  }

  updateMeshes() {
    for (let i = 0; i < this.objects.length; i++) {
      const obj = this.objects[i];
      obj.updateFromSource();

      if (obj.mesh) {
        // Update position with bouncy animation
        obj.mesh.position.set(
          obj.position[0],
          obj.position[1] + Math.sin(this.i + obj.mesh.userData.phase) * 0.3,
          obj.position[2]
        );

        // Scale animation
        const scale = 1 + Math.sin(this.i * 2 + obj.mesh.userData.phase) * 0.1;
        obj.mesh.scale.set(scale, scale, scale);

        // Color updates
        if (obj.mesh.material.color.getHex() !== obj.color) {
          obj.mesh.material.color.set(obj.color);
          obj.mesh.material.emissive.set(obj.color);
        }

        // Size updates (simplified)
        const newSize = this.calculateVisualSize(obj.size);
        if (obj.mesh.children[0].geometry.parameters.radius !== newSize * 1.1) {
          obj.mesh.geometry.dispose();
          obj.mesh.geometry = new THREE.SphereGeometry(
            newSize,
            this.sphereSegments,
            this.sphereSegments
          );
          obj.mesh.children[0].geometry.dispose();
          obj.mesh.children[0].geometry = new THREE.SphereGeometry(
            newSize * 1.1,
            this.sphereSegments,
            this.sphereSegments
          );
        }
      }
    }
  }

  // Handle window resize
  onWindowResize() {
    // this.camera.aspect = window.innerWidth / window.innerHeight;
    // this.camera.updateProjectionMatrix();
    // this.renderer.setSize(window.innerWidth, window.innerHeight);
  }

  animate() {
    requestAnimationFrame(this.animate.bind(this));

    // Custom animation with more exaggerated motion
    if (this.customAnimationCallback) {
      this.customAnimationCallback(this.objects, this.i);
    } else {
      this.objects.forEach((obj, index) => {
        if (obj.sourceObject) return;
        const radius = 3 + index * 2;
        const speed = 0.02 / (1 + index * 0.2);
        const angle = this.i * speed;

        obj.updatePosition([
          radius * Math.sin(angle),
          0,
          radius * Math.cos(angle)
        ]);
      });
    }

    this.updateMeshes();
    this.i += 0.02;
    this.renderer.render(this.scene, this.camera);
  }

  // Clear all objects
  clear() {
    for (const obj of this.objects) {
      if (obj.mesh) {
        this.scene.remove(obj.mesh);
        obj.mesh.geometry.dispose();
        obj.mesh.material.dispose();
      }
    }
    this.objects = [];
  }
}

export default SolarSystem;
