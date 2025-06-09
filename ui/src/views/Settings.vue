<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useProfileStore } from '../stores/profile'
import { useAdminStore } from '../stores/admin'
import AdminUsers from '../components/admin/AdminUsers.vue'
import AdminGroups from '../components/admin/AdminGroups.vue'
import AdminSpaces from '../components/admin/AdminSpaces.vue'
import { UserIcon, UsersIcon, UserGroupIcon, Square3Stack3DIcon } from '@heroicons/vue/24/outline';

const profileStore = useProfileStore()
const adminStore = useAdminStore()

const activeTab = ref('profile')

const isAdmin = computed(() => profileStore.profile?.is_admin ?? false)

const tabs = computed(() => {
  const baseTabs = [
    { id: 'profile', label: 'Profile', icon: UserIcon }
  ]
  
  if (isAdmin.value) {
    baseTabs.push(
      { id: 'users', label: 'Users', icon: UsersIcon },
      { id: 'groups', label: 'Groups', icon: UserGroupIcon },
      { id: 'spaces', label: 'Spaces', icon: Square3Stack3DIcon }
    )
  }
  
  return baseTabs
})

onMounted(async () => {
  await profileStore.fetchProfile()
})

const switchTab = (tabId: string) => {
  activeTab.value = tabId
  
  // Load admin data when switching to admin tabs
  if (isAdmin.value) {
    if (tabId === 'users') {
      adminStore.fetchUsers()
    } else if (tabId === 'groups') {
      adminStore.fetchGroups()
    } else if (tabId === 'spaces') {
      adminStore.fetchSpaces()
    }
  }
}
</script>

<template>
  <main class="flex-1 flex flex-col overflow-hidden bg-white">
    <div class="flex flex-col h-full">
      <div class="sticky top-0 border-b border-e border-gray-200 bg-white">
        <div class="flex justify-between items-center h-16 px-8">
          <h1 class="w-full text-2xl font-medium text-gray-600 bg-transparent focus:outline-none text-center">Settings</h1>
        </div>
      </div>
      <div class="sticky top-0 border-b border-e border-gray-200 bg-white">
        <div v-if="profileStore.loading" class="flex justify-center items-center h-32">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>
        
        <!-- Settings Content -->
        <div v-else class="bg-white">
          <!-- Tab Navigation -->
          <div class="border-b border-gray-200">
            <nav class="-mb-px flex space-x-8 justify-center">
              <button
                v-for="tab in tabs"
                :key="tab.id"
                @click="switchTab(tab.id)"
                :class="[
                  'py-4 px-1 border-b-2 font-medium text-sm whitespace-nowrap transition-colors duration-200 flex items-center',
                  activeTab === tab.id
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                <component :is="tab.icon" class="w-5 h-5 mr-2" />
                {{ tab.label }}
              </button>
            </nav>
          </div>
        </div>
        <div class="mx-auto">
          
            <!-- Profile Tab -->
            <div v-if="activeTab === 'profile'">
              <div class="max-w-lg mx-auto">                
                <div v-if="profileStore.profile" class="space-y-4 text-center">
                    <div class="grid grid-cols-1 gap-4">
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
                        <p class="text-gray-900">{{ profileStore.profile.email }}</p>
                      </div>
                      
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Name</label>
                        <p class="text-gray-900">{{ profileStore.profile.name || 'Not set' }}</p>
                      </div>
                      
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Role</label>
                        <span :class="[
                          'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                          profileStore.profile.is_admin 
                            ? 'bg-purple-100 text-purple-800' 
                            : 'bg-gray-100 text-gray-800'
                        ]">
                          {{ profileStore.profile.is_admin ? 'Administrator' : 'User' }}
                        </span>
                      </div>
                      
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Created</label>
                        <p class="text-gray-900">
                          {{ profileStore.profile.created_at }}
                        </p>
                      </div>
                    </div>
                </div>
              </div>
            </div>
            
            <!-- Admin Users Tab -->
            <div v-else-if="activeTab === 'users' && isAdmin">
              <AdminUsers />
            </div>
            
            <!-- Admin Groups Tab -->
            <div v-else-if="activeTab === 'groups' && isAdmin">
              <AdminGroups />
            </div>
            
            <!-- Admin Spaces Tab -->
            <div v-else-if="activeTab === 'spaces' && isAdmin">
              <AdminSpaces />
            </div>
        </div>
      </div>
    </div>
  </main>
</template>