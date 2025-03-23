<script setup>
import { ref, onMounted, nextTick, watch } from 'vue';
import SolarSystem from './plot.js';
import Planet from './planet.js';

const enablePrediction = ref(false);
const spaceElement = ref(null);
let solarSystem;
onMounted(async () => {
  await nextTick(); // Wait for DOM updates
  if (!spaceElement.value) return;
  
  solarSystem = new SolarSystem(0, [], {
    minSize: 0.1,
    maxSize: 1.0,
    sizeScaleFactor: 0.1,
    sphereSegments: 24,
    x: 720,
    y: 720,
    element: '.space'
  });
  
  const G = 1.4881851e-10; // Gravitational constant in m³/kg/s²
  const sunMass = 1.989e6; // Mass of the Sun in kg
  
  // Calculate orbital velocity in AU/day
  const calculateOrbitalVelocityAUPerDay = (distanceAU) => {
    const velocityMS = Math.sqrt(G * sunMass / distanceAU); // Velocity in m/s
    return velocityMS;
  };
  
  console.log(calculateOrbitalVelocityAUPerDay(1));
  
  const planets = [
    // Name, mass, diameter, color, initial position [x,y,z], initial velocity [x,y,z]
    new Planet("Sun", 1989000, 0.05, 0xffff00, [0, 0, 0], [0, 0, 0]),
    new Planet("Mercury", 0.330, 0.05, 0xaaaaaa, [0.39, 0, 0], [0, 0, calculateOrbitalVelocityAUPerDay(0.39)]),
    new Planet("Venus", 4.87, 0.05, 0xe39e1c, [0.72, 0, 0], [0, 0, calculateOrbitalVelocityAUPerDay(0.72)]),
    new Planet("Earth", 5.97, 0.05, 0x2233dd, [1.0, 0, 0], [0, 0, calculateOrbitalVelocityAUPerDay(1.0)]),
    new Planet("Mars", 0.642, 0.05, 0xdd3333, [1.52, 0, 0], [0, 0, calculateOrbitalVelocityAUPerDay(1.52)]),
    // new Planet("Jupiter", 1898, 0.1, 0xdda52c, [5.2, 0, 0], [0, 0, calculateOrbitalVelocityAUPerDay(5.2)])
  ];
  
  for (const planet of planets) {
    solarSystem.addObject(planet);
  }
  
  solarSystem.setAnimationCallback(() => {
    for (const planet of planets) {
      planet.update(0.5);
    }
  });
  
  solarSystem.animate();
});

function togglePrediction(){
  enablePrediction.value = !enablePrediction.value;
  solarSystem.enablePrediction(enablePrediction.value);
}
</script>

<template>
  <div ref="spaceElement" class="space" />
  <button class="enable-button" @click="togglePrediction">{{enablePrediction?'Disable':'Enable'}} Prediction</button>
</template>

<style scoped>
.enable-button {
  background-color: rgba(65, 90, 160, 0.3);
  border: 1px solid rgba(65, 120, 255, 0.2);
  border-top-left-radius: 0;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #a0c0ff;
  cursor: pointer;
  transition: all 0.2s;
}

.enable-button:hover {
  background-color: rgba(65, 90, 160, 0.5);
}

</style>
