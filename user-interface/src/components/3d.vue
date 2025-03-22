<script>
import SolarSystem from "./plot.js";
import { onMounted, ref, nextTick } from "vue";

class Planet {
  constructor(name, diameter, orbitRadius, rotationSpeed, color) {
    this.name = name;
    this.diameter = diameter; // Actual size
    this.orbitRadius = orbitRadius;
    this.rotationSpeed = rotationSpeed;
    this.color = color;
    this.angle = 0;
  }

  update(deltaTime) {
    this.angle += this.rotationSpeed * deltaTime;
  }

  getPosition() {
    return [
      this.orbitRadius * Math.sin(this.angle),
      0,
      this.orbitRadius * Math.cos(this.angle)
    ];
  }

  getColor() {
    return this.color;
  }

  getSize() {
    return this.diameter;
  }
}

export default {
  setup() {
    const spaceElement = ref(null);

    onMounted(async () => {
      await nextTick(); // Wait for DOM updates

      if (!spaceElement.value) return;

      const solarSystem = new SolarSystem(0, [], {
        minSize: 0.1,
        maxSize: 1.0,
        sizeScaleFactor: 0.1,
        sphereSegments: 24,
        x: 720,
        y: 720,
        element: ".space"
      });

      const planets = [
        new Planet("Mercury", 0.038, 1.8, 0.02, 0xaaaaaa),
        new Planet("Venus", 0.095, 3.0, 0.015, 0xe39e1c),
        new Planet("Earth", 0.1, 5.0, 0.01, 0x2233dd),
        new Planet("Mars", 0.053, 7.0, 0.008, 0xdd3333),
        new Planet("Jupiter", 0.11, 12.0, 0.002, 0xdda52c)
      ];

      for (const planet of planets) {
        solarSystem.addObject(planet.name, planet);
      }

      solarSystem.setAnimationCallback(() => {
        for (const planet of planets) {
          planet.update(0.5);
        }
      });

      solarSystem.animate();
    });

    return { spaceElement };
  }
};
</script>

<template>
  <div ref="spaceElement" class="space" />
</template>

<style scoped></style>
