<template>
  <div class="cubespace-container">
    <nav class="nav-bar">
      <!-- Logo/Brand -->
      <div class="brand">
        <span class="brand-name">CubeSpace</span>
      </div>

      <!-- Navigation Links -->
      <div class="nav-links">
        <a href="#" class="nav-link">Dashboard</a>
        <a href="#" class="nav-link">Tracking</a>
        <a href="#" class="nav-link">Analytics</a>
        <a href="#" class="nav-link">Settings</a>
      </div>

      <!-- Action Icons -->
      <div class="action-icons">
        <button class="icon-button">
          <Search class="icon" />
        </button>
        <button class="icon-button">
          <Bell class="icon" />
        </button>
        <button class="icon-button">
          <User class="icon" />
        </button>
        <button class="icon-button mobile-menu">
          <Menu class="icon" />
        </button>
      </div>
    </nav>
  </div>
</template>

<script setup>
import { Search, Bell, User, Menu } from "lucide-vue-next";
</script>

<style scoped>
.cubespace-container {
  font-family: "Arial", sans-serif;
  background-color: #0a0a1a;
  background-image: radial-gradient(
    circle at 50% 50%,
    rgba(16, 24, 48, 0.8) 0%,
    rgba(8, 8, 24, 1) 100%
  );
  min-height: 100vh;
}

.nav-bar {
  width: 100%;
  background-color: rgba(16, 24, 48, 0.4);
  border-bottom: 1px solid rgba(65, 90, 160, 0.3);
  padding: 12px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

.brand {
  display: flex;
  align-items: center;
}

.brand-name {
  color: #a0c0ff;
  font-weight: bold;
  font-size: 1.25rem;
  letter-spacing: 0.5px;
}

.nav-links {
  display: flex;
  gap: 24px;
}

.nav-link {
  color: #a0c0ff;
  font-size: 0.875rem;
  text-decoration: none;
  transition: color 0.2s;
}

.nav-link:hover {
  color: #ffffff;
}

.action-icons {
  display: flex;
  gap: 12px;
}

.icon-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(65, 90, 160, 0.3);
  border: 1px solid rgba(65, 120, 255, 0.3);
  border-radius: 50%;
  width: 32px;
  height: 32px;
  color: #a0c0ff;
  cursor: pointer;
  transition: background-color 0.2s;
}

.icon-button:hover {
  background-color: rgba(65, 90, 160, 0.5);
}

.icon {
  width: 16px;
  height: 16px;
}

/* Responsive styles */
@media (max-width: 768px) {
  .nav-links {
    display: none;
  }
}

@media (min-width: 769px) {
  .mobile-menu {
    display: none;
  }
}
</style>