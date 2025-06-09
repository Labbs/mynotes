export interface UserPreferences {
  ui: {
    expanded_documents?: string[]
    expanded_spaces?: string[]
    sidebarCollapsed?: boolean
    sidebarWidth?: number
  }
}

export interface Space {
  id: string
  name: string
  description?: string
  slug?: string
  icon?: string
  icon_color?: string
  members?: Member[]
  members_with_users_or_groups?: MemberWithUserOrGroup[]
  created_at?: string
  updated_at?: string
}

export interface Favorite {
  id: string
  user_id: string
  document_id?: string
  database_id?: string
  position?: string
  document?: Document
  created_at?: string
}

export interface User {
  id: string
  email: string
  name: string
  avatar?: string
  is_admin?: boolean
  created_at?: string
  updated_at?: string
  groups?: Group[]
}

export interface Document {
  id: string
  name: string
  space_id: string
  type: 'document' | 'excalidraw' | 'database'
  parent_id?: string
  slug?: string
  // properties: Property[]
  config: Config
  public?: boolean
  members?: Member[]
  content?: string
  created_at?: string
  updated_at?: string
}

export interface Config {
  full_width?: boolean
  icon?: string
  lock?: boolean
  header_background?: string
}

export interface Group {
  id: string
  name: string
  description?: string
  role: string
  created_at?: string
  updated_at?: string
  users?: User[]
}

export interface Member {
  id: string
  type: 'user' | 'group'
  access: 'viewer' | 'editor' | 'comment' | 'full'
}

export interface MemberWithUserOrGroup extends Member {
  user?: User
  group?: Group
}