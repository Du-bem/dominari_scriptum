<template>
  <div class="astral-tracker-container">
    <!-- Astral Body Component with Dropdown -->
    <div class="tracker-panel">
      <div class="panel-header">
        <h2 class="text-lg font-bold text-white">Viewer</h2>

        <!-- Astral Body Type Dropdown -->
        <div class="astral-body-dropdown">
          <button @click="dropdownOpen = !dropdownOpen" class="dropdown-button">
            {{ selectedBodyType }}
            <ChevronDown class="h-4 w-4 ml-1" />
          </button>

          <div v-if="dropdownOpen" class="dropdown-menu">
            <div
              v-for="type in bodyTypes"
              :key="type"
              @click="selectBodyType(type)"
              class="dropdown-item"
            >
              {{ type }}
            </div>
          </div>
        </div>
      </div>

      <!-- Astral Body Data Card -->
      <div class="astral-data-card">
        <!-- Loading State -->
        <div v-if="loading" class="loading-state">
          <Loader class="animate-spin h-6 w-6 text-blue-500" />
          <p class="mt-2 text-gray-300 text-sm">Loading data...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="error-state">
          <AlertCircle class="h-8 w-8 text-red-500 mb-2" />
          <p class="text-red-400 text-sm font-semibold">Error loading data</p>
          <p class="text-gray-400 mt-1 text-xs">{{ error }}</p>
          <button
            @click="fetchData"
            class="mt-2 bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 text-xs rounded-md transition-colors"
          >
            Retry
          </button>
        </div>

        <!-- Data Display -->
        <div v-else-if="astralData" class="astral-content">
          <!-- Object Type Icon and Name -->
          <div class="object-header">
            <component
              :is="getObjectIcon(astralData.name)"
              class="h-6 w-6 text-blue-400"
            />
            <h3 class="text-lg font-bold text-white ml-2">
              {{ astralData.name }}
            </h3>
            <div class="classification-badge ml-auto">
              {{ classifyObject(astralData.name) }}
            </div>
          </div>

          <!-- Time Stamp -->
          <div class="time-stamp">
            <Clock class="inline-block h-3 w-3 mr-1 text-gray-400" />
            <span class="text-gray-300 text-xs">{{
              formatTime(astralData.time)
            }}</span>
          </div>

          <!-- Position Data -->
          <div class="data-section">
            <div class="section-header">
              <MapPin class="h-4 w-4 mr-1 text-blue-400" />
              <h4 class="text-sm font-semibold text-blue-400">Position</h4>
            </div>
            <div class="coordinate-grid">
              <div class="coordinate">
                <span class="coordinate-label">X</span>
                <span class="coordinate-value">{{
                  formatCoordinate(astralData.position[0])
                }}</span>
              </div>
              <div class="coordinate">
                <span class="coordinate-label">Y</span>
                <span class="coordinate-value">{{
                  formatCoordinate(astralData.position[1])
                }}</span>
              </div>
              <div class="coordinate">
                <span class="coordinate-label">Z</span>
                <span class="coordinate-value">{{
                  formatCoordinate(astralData.position[2])
                }}</span>
              </div>
            </div>
          </div>

          <!-- Velocity Data -->
          <div class="data-section">
            <div class="section-header">
              <Zap class="h-4 w-4 mr-1 text-green-400" />
              <h4 class="text-sm font-semibold text-green-400">Velocity</h4>
            </div>
            <div class="coordinate-grid">
              <div class="coordinate">
                <span class="coordinate-label">X</span>
                <span class="coordinate-value">{{
                  formatVelocity(astralData.velocity[0])
                }}</span>
              </div>
              <div class="coordinate">
                <span class="coordinate-label">Y</span>
                <span class="coordinate-value">{{
                  formatVelocity(astralData.velocity[1])
                }}</span>
              </div>
              <div class="coordinate">
                <span class="coordinate-label">Z</span>
                <span class="coordinate-value">{{
                  formatVelocity(astralData.velocity[2])
                }}</span>
              </div>
            </div>
          </div>

          <!-- Hash Data -->
          <div class="data-section">
            <div class="section-header checksum">
              <Shield class="h-4 w-4 mr-1 text-green-400" />
              <h4 class="text-sm font-semibold text-green-400">Hash</h4>
              <div class="checksum-data">
                {{ formatChecksum(astralData.checksum) }}
              </div>
            </div>
          </div>
        </div>

        <!-- No Data State -->
        <div v-else class="no-data-state">
          <Database class="h-8 w-8 text-gray-600 mb-2" />
          <p class="text-gray-400 text-sm">No data available</p>
          <button
            @click="fetchData"
            class="mt-2 bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 text-xs rounded-md transition-colors"
          >
            Fetch Data
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import {
  Loader,
  AlertCircle,
  Database,
  Clock,
  MapPin,
  Zap,
  Star,
  Globe,
  Activity,
  Moon,
  ChevronDown,
  RefreshCw,
  Shield
} from "lucide-vue-next";
import {getDataFromAPI} from "./getData.js";

// Astral Body Types Dropdown
const dropdownOpen = ref(false);
const bodyTypes = ref([
  "All Bodies",
  "Asteroids",
  "Planets",
  "Stars",
  "Moons",
  "Satellites"
]);
const selectedBodyType = ref("All Bodies");

const selectBodyType = type => {
  selectedBodyType.value = type;
  dropdownOpen.value = false;
  fetchData(); // Fetch data based on selected type
};

// API Data
const astralData = ref(null);
const loading = ref(false);
const error = ref(null);

// Format coordinate to 2 decimal places with unit
const formatCoordinate = value => {
  if (value === undefined || value === null) return "N/A";
  return `${value.toLocaleString(undefined, { maximumFractionDigits: 2 })} km`;
};

const formatChecksum = value => {
  if (value.length > 22) {
    return `${value.slice(0, 22)}...`;
  }
  return value;
};

// Format velocity to 2 decimal places with unit
const formatVelocity = value => {
  if (value === undefined || value === null) return "N/A";
  return `${value.toLocaleString(undefined, { maximumFractionDigits: 2 })} m/s`;
};

// Calculate magnitude of a vector
const calculateMagnitude = vector => {
  if (!vector || !Array.isArray(vector) || vector.length < 3) return "N/A";

  const sumOfSquares = vector.reduce(
    (sum, component) => sum + Math.pow(component, 2),
    0
  );
  const magnitude = Math.sqrt(sumOfSquares);

  return magnitude.toLocaleString(undefined, { maximumFractionDigits: 2 });
};

// Format time from dd/mm/yr 00:00:00 to a more readable format
const formatTime = timeString => {
  if (!timeString) return "Unknown time";

  try {
    // Parse the date parts
    const parts = timeString.split(" ");
    const datePart = parts[0];
    const timePart = parts[1] || "00:00:00";

    const [day, month, year] = datePart.split("/");

    // Create a date object
    const date = new Date(`${year}-${month}-${day}T${timePart}`);

    // Format the date
    return date.toLocaleString("en-US", {
      day: "numeric",
      month: "short",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit"
    });
  } catch (err) {
    console.error("Error formatting time:", err);
    return timeString; // Return original if parsing fails
  }
};

// Determine icon based on object name
const getObjectIcon = name => {
  const nameLower = name.toLowerCase();

  if (nameLower.includes("asteroid")) return Activity;
  if (nameLower.includes("planet")) return Globe;
  if (nameLower.includes("star")) return Star;
  if (nameLower.includes("moon") || nameLower.includes("satellite"))
    return Moon;

  // Default icon
  return Star;
};

// Classify object based on name
const classifyObject = name => {
  const nameLower = name.toLowerCase();

  if (nameLower.includes("asteroid")) return "Asteroid";
  if (nameLower.includes("planet")) return "Planet";
  if (nameLower.includes("star")) return "Star";
  if (nameLower.includes("moon")) return "Moon";
  if (nameLower.includes("satellite")) return "Satellite";

  return "Unknown";
};

// Fetch data based on selected body type
const fetchData = async () => {
  const data = await getDataFromAPI();
  loading.value = true;
  error.value = null;

  try {
    // In a real application, this would filter by the selected body type
    // For demonstration, we'll simulate a delay
    // Generate data based on selected type
    const type =
      selectedBodyType.value === "All Bodies"
        ? bodyTypes.value[
            Math.floor(Math.random() * (bodyTypes.value.length - 1)) + 1
          ]
        : selectedBodyType.value;

    // Remove 's' from the end if present (e.g., "Asteroids" -> "Asteroid")
    const singularType = type.endsWith("s") ? type.slice(0, -1) : type;

    const id = data.name;

    const now = new Date(data.time);
    const day = now
      .getDate()
      .toString()
      .padStart(2, "0");
    const month = (now.getMonth() + 1).toString().padStart(2, "0");
    const year = now.getFullYear();
    const hours = now
      .getHours()
      .toString()
      .padStart(2, "0");
    const minutes = now
      .getMinutes()
      .toString()
      .padStart(2, "0");
    const seconds = now
      .getSeconds()
      .toString()
      .padStart(2, "0");

    astralData.value = {
      name: `${singularType} ${id}`,
      time: `${day}/${month}/${year} ${hours}:${minutes}:${seconds}`,
      position: [data.position[0], data.position[1], data.position[2]],
      velocity: [data.velocity[0], data.velocity[1], data.velocity[2]],
      checksum: data.checksum
    };
  } catch (err) {
    console.error("Error fetching data:", err);
    error.value = err.message || "Failed to fetch astral body data";
  } finally {
    loading.value = false;
  }
};

// Fetch data on component mount
onMounted(() => {
  fetchData();

  // Close dropdown when clicking outside
  document.addEventListener("click", e => {
    const dropdown = document.querySelector(".astral-body-dropdown");
    if (dropdown && !dropdown.contains(e.target)) {
      dropdownOpen.value = false;
    }
  });
});
</script>

<style scoped>
.astral-tracker-container {
  font-family: "Arial", sans-serif;
  display: flex;
  justify-content: center;
  background-color: #0a0a1a;
  background-image: radial-gradient(
    circle at 50% 50%,
    rgba(16, 24, 48, 0.8) 0%,
    rgba(8, 8, 24, 1) 100%
  );
}

/* Tracker Panel Styling */
.tracker-panel {
  width: 100%;
  background-color: rgba(16, 24, 48, 0.4);
  border-top-left-radius: 12px;
  border: 1px solid rgba(65, 90, 160, 0.3);
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.5), 0 0 20px rgba(65, 120, 255, 0.1);
}

.panel-header {
  background-color: rgba(0, 0, 0, 0.3);
  padding: 12px 16px;
  border-bottom: 1px solid rgba(65, 90, 160, 0.3);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Astral Body Dropdown */
.astral-body-dropdown {
  position: relative;
}

.dropdown-button {
  display: flex;
  align-items: center;
  background-color: rgba(65, 90, 160, 0.3);
  border: 1px solid rgba(65, 120, 255, 0.3);
  border-radius: 4px;
  padding: 4px 10px;
  font-size: 0.8rem;
  color: #a0c0ff;
  cursor: pointer;
  transition: all 0.2s;
}

.dropdown-button:hover {
  background-color: rgba(65, 90, 160, 0.5);
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  background-color: rgba(20, 30, 60, 0.95);
  border: 1px solid rgba(65, 120, 255, 0.3);
  border-radius: 4px;
  width: 150px;
  z-index: 10;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.dropdown-item {
  padding: 8px 12px;
  font-size: 0.8rem;
  cursor: pointer;
  transition: background-color 0.2s;
  color: #fff;
}

.dropdown-item:hover {
  background-color: rgba(65, 90, 160, 0.5);
}

/* Astral Data Card */
.astral-data-card {
  padding: 16px;
  height: 400px;
  display: flex;
  flex-direction: column;
}

.loading-state,
.error-state,
.no-data-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
  color: #fff;
}

.astral-content {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.object-header {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
  margin-left: 8px;
  margin-right: 8px;
  gap: 4px;
}

.classification-badge {
  background-color: rgba(65, 120, 255, 0.2);
  color: #64b5f6;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.7rem;
  border: 1px solid rgba(65, 120, 255, 0.4);
  margin-left: auto;
}

.time-stamp {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-bottom: 16px;
  color: #aaa;
  font-size: 0.8rem;
}

.data-section {
  margin-bottom: 16px;
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 6px;
  padding: 12px;
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.section-header {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
  gap: 4px;
}

.coordinate-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
  margin-bottom: 8px;
}

.coordinate {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.3);
  padding: 6px;
  border-radius: 4px;
}

.coordinate-label {
  font-weight: bold;
  color: #aaa;
  font-size: 0.7rem;
  margin-bottom: 2px;
}

.coordinate-value {
  font-family: "Courier New", monospace;
  font-size: 0.7rem;
  color: #fff;
}

.velocity-magnitude {
  margin-top: 8px;
  text-align: right;
  padding: 4px 8px;
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 4px;
}

.last-updated {
  margin-top: auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.refresh-button {
  background-color: rgba(65, 90, 160, 0.3);
  border: 1px solid rgba(65, 120, 255, 0.2);
  border-radius: 50%;
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #a0c0ff;
  cursor: pointer;
  transition: all 0.2s;
}

.refresh-button:hover {
  background-color: rgba(65, 90, 160, 0.5);
}

.checksum {
  display: flex;
  align-items: center;
}

.checksum-data {
  margin-left: auto;
}

@media (max-width: 400px) {
  .tracker-panel {
    width: 100%;
  }

  .coordinate-grid {
    grid-template-columns: 1fr;
  }
}
</style>
