import { defineStore } from 'pinia';
import { meApi } from '../api/me';
import type { Favorite } from '../api/interface';

interface FavoriteState {
  favorites: Favorite[];
  loading: boolean;
  error: string | null;
}

export const useFavoriteStore = defineStore('favorite', {
  state: (): FavoriteState => ({
    favorites: [],
    loading: false,
    error: null,
  }),

  actions: {
    async fetchFavorites() {
      this.loading = true;
      this.error = null;

      try {
        const { data } = await meApi.getMyFavorites();
        this.favorites = Array.isArray(data) ? data : data.favorites || [];
      } catch (error) {
        this.error = 'Failed to fetch favorites';
        console.error('Error fetching favorites:', error);
      } finally {
        this.loading = false;
      }
    },

    async addFavorite(documentId: string) {
      try {
        const { data } = await meApi.addFavorite(documentId);
        this.favorites = data;
        return data;
      } catch (error) {
        console.error('Error adding favorite:', error);
        throw error;
      }
    },

    async unFavorite(documentId: string) {
      try {
        const { data } = await meApi.unFavorite(documentId);
        this.favorites = data;
      } catch (error) {
        console.error('Error removing favorite:', error);
        throw error;
      }
    },
  }
});