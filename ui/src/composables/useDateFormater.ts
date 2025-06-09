export function useDateFormatter() {
  const formatDate = (dateString: string | null | undefined, locale = 'fr-FR') => {
    if (!dateString) return 'N/A'
    return new Date(dateString).toLocaleString(locale, { 
      year: 'numeric', 
      month: 'short', 
      day: 'numeric', 
      hour: '2-digit', 
      minute: '2-digit' 
    })
  }

  return { formatDate }
}