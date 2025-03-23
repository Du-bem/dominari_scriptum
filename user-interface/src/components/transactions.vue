<template>
    <div class="astral-tracker-container">
      <div class="tracker-panel">
        <div class="panel-header">
          <h2 style="margin: 0; color: #a0c0ff; font-size: 1.2rem;">TRANSACTIONS BALANCE</h2>
          <div class="astral-body-dropdown">
            <button class="dropdown-button">
              <span>Filter</span>
              <span style="margin-left: 4px;">▼</span>
            </button>
          </div>
        </div>
        
        <div class="astral-data-card">
          <div v-if="loading" class="loading-state">
            <p>Loading transactions...</p>
          </div>
          <div v-else-if="error" class="error-state">
            <p>Error loading transactions</p>
          </div>
          <div v-else-if="transactions.length === 0" class="no-data-state">
            <p>No transactions found</p>
          </div>
          <div v-else class="astral-content">
            <div class="data-section">
              <div class="section-header">
                <span style="color: #a0c0ff; font-size: 0.9rem;">Transactions</span>
              </div>
              <div class="transaction-list">
                <div 
                  v-for="(transaction, index) in transactions" 
                  :key="index"
                  class="transaction-item"
                  style="margin-bottom: 12px; padding: 8px; background-color: rgba(0, 0, 0, 0.3); border-radius: 4px; display: flex; justify-content: space-between; align-items: center;"
                >
                  <div class="transaction-hash" style="color: #64b5f6; font-family: 'Courier New', monospace; font-size: 0.8rem; cursor: pointer;">
                    LINK HASH {{ transaction.id }}
                  </div>
                  <div class="transaction-amount" style="color: #fff; font-size: 0.8rem;">
                    {{ transaction.amount }}
                  </div>
                </div>
              </div>
            </div>
            
            <div class="data-section">
              <div class="section-header">
                <span style="color: #a0c0ff; font-size: 0.9rem;">Balance Summary</span>
              </div>
              <div 
                style="display: flex; justify-content: space-between; align-items: center; padding: 8px; background-color: rgba(0, 0, 0, 0.3); border-radius: 4px;"
              >
                <span style="color: #aaa; font-size: 0.8rem;">Total Balance:</span>
                <span style="color: #fff; font-size: 0.9rem; font-weight: bold;">{{ totalBalance }}</span>
              </div>
            </div>
            
            <div class="last-updated">
              <span style="color: #aaa; font-size: 0.7rem;">Last updated: {{ lastUpdated }}</span>
              <button class="refresh-button" @click="refreshData">
                ↻
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import {getChainFromAPI} from "./getData.js";
  export default {
    name: 'TransactionsBalance',
    data() {
      return {
        loading: false,
        error: null,
        transactions: [
          { id: '1', hash: '0x8a23f12d4c78e90bf123a', amount: '0.0456 ETH' },
          { id: '2', hash: '0x3e67d89f1a45c78e20bf9', amount: '0.0123 ETH' },
          { id: '3', hash: '0x7b92e45d1c8a903fb578', amount: '0.0789 ETH' },
          { id: '4', hash: '0x2c46d90f8a71e25b3d4c', amount: '0.0321 ETH' }
        ],
        lastUpdated: new Date().toLocaleTimeString()
      }
    },
    computed: {
      totalBalance() {
        // This is a mock calculation - in a real app, you would calculate based on actual transaction values
        return '0.1689 ETH'
      }
    },
    methods: {
      refreshData() {
        this.loading = true;
        getChainFromAPI().then(()=>{
            this.loading = false;
            this.lastUpdated = new Date().toLocaleTimeString();
        });
        
      }
    },
    mounted() {
      // Initial data load
      this.refreshData()
    }
  }
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
    border-radius: 12px;
    border: 1px solid rgba(65, 90, 160, 0.3);
    overflow: hidden;
    margin-top: 24px;
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
  </style>