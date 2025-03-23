class Planet {
  constructor(name, mass, diameter, color, pos, vel) {
    this.name = name;
    this.mass = mass;
    this.diameter = diameter;
    this.color = color;
    this.pos = pos; // [x, y, z]
    this.vel = vel; // [x, y, z]
  }

  update(deltaTime, allPlanets) {
    const G = 1.4881851e-10; // Gravitational constant
    
    const myIndex = allPlanets.findIndex(planet => planet.name === this.name);
    
    const calculateGravitationalForce = (planet1, planet2) => {
      const rVector = [
        planet2.pos[0] - planet1.pos[0],
        planet2.pos[1] - planet1.pos[1],
        planet2.pos[2] - planet1.pos[2]
      ];
      
      const distanceSquared = rVector[0]**2 + rVector[1]**2 + rVector[2]**2;
      const distance = Math.sqrt(distanceSquared);
      
      const forceMagnitude = G * planet1.mass * planet2.mass / distanceSquared;
      
      return [
        forceMagnitude * rVector[0] / distance,
        forceMagnitude * rVector[1] / distance,
        forceMagnitude * rVector[2] / distance
      ];
    };
    
    const derivatives = (position, velocity) => {
      let netForce = [0, 0, 0];
      
      // Corrected loop to include all other planets
      for (let i = 0; i < allPlanets.length; i++) {
        if (i !== myIndex) { // Changed from 'i > myIndex' to 'i !== myIndex'
          const otherPlanet = allPlanets[i];
          
          const planet1 = {
            pos: position,
            mass: this.mass
          };
          
          const force = calculateGravitationalForce(planet1, otherPlanet);
          netForce[0] += force[0];
          netForce[1] += force[1];
          netForce[2] += force[2];
        }
      }
      
      const dPos = [...velocity];
      const dVel = [
        netForce[0] / this.mass,
        netForce[1] / this.mass,
        netForce[2] / this.mass
      ];
      
      return { dPos, dVel };
    };
    
    // RK4 implementation remains the same but has limitations (see note below)
    const k1 = derivatives(this.pos, this.vel);
    
    const tempPos2 = [
      this.pos[0] + 0.5 * k1.dPos[0] * deltaTime,
      this.pos[1] + 0.5 * k1.dPos[1] * deltaTime,
      this.pos[2] + 0.5 * k1.dPos[2] * deltaTime
    ];
    const tempVel2 = [
      this.vel[0] + 0.5 * k1.dVel[0] * deltaTime,
      this.vel[1] + 0.5 * k1.dVel[1] * deltaTime,
      this.vel[2] + 0.5 * k1.dVel[2] * deltaTime
    ];
    const k2 = derivatives(tempPos2, tempVel2);
    
    const tempPos3 = [
      this.pos[0] + 0.5 * k2.dPos[0] * deltaTime,
      this.pos[1] + 0.5 * k2.dPos[1] * deltaTime,
      this.pos[2] + 0.5 * k2.dPos[2] * deltaTime
    ];
    const tempVel3 = [
      this.vel[0] + 0.5 * k2.dVel[0] * deltaTime,
      this.vel[1] + 0.5 * k2.dVel[1] * deltaTime,
      this.vel[2] + 0.5 * k2.dVel[2] * deltaTime
    ];
    const k3 = derivatives(tempPos3, tempVel3);
    
    const tempPos4 = [
      this.pos[0] + k3.dPos[0] * deltaTime,
      this.pos[1] + k3.dPos[1] * deltaTime,
      this.pos[2] + k3.dPos[2] * deltaTime
    ];
    const tempVel4 = [
      this.vel[0] + k3.dVel[0] * deltaTime,
      this.vel[1] + k3.dVel[1] * deltaTime,
      this.vel[2] + k3.dVel[2] * deltaTime
    ];
    const k4 = derivatives(tempPos4, tempVel4);
    
    this.pos[0] += (k1.dPos[0] + 2 * k2.dPos[0] + 2 * k3.dPos[0] + k4.dPos[0]) * deltaTime / 6;
    this.pos[1] += (k1.dPos[1] + 2 * k2.dPos[1] + 2 * k3.dPos[1] + k4.dPos[1]) * deltaTime / 6;
    this.pos[2] += (k1.dPos[2] + 2 * k2.dPos[2] + 2 * k3.dPos[2] + k4.dPos[2]) * deltaTime / 6;
    
    this.vel[0] += (k1.dVel[0] + 2 * k2.dVel[0] + 2 * k3.dVel[0] + k4.dVel[0]) * deltaTime / 6;
    this.vel[1] += (k1.dVel[1] + 2 * k2.dVel[1] + 2 * k3.dVel[1] + k4.dVel[1]) * deltaTime / 6;
    this.vel[2] += (k1.dVel[2] + 2 * k2.dVel[2] + 2 * k3.dVel[2] + k4.dVel[2]) * deltaTime / 6;
  }
  

  getPosition() {
    return this.pos;
  }
  
  getVelocity() {
    return this.vel;
  }

  getColor() {
    return this.color;
  }

  getSize() {
    return this.diameter;
  }
}

export default Planet;