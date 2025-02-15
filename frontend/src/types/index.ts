
export interface UserProp {
  id: string;
  firstname: string;
  lastname: string;
  email: string;
  street: string;
  city: string;
  state: string;
  zipcode: string;
  createdAt: string;
}

export interface PaginationProp {
  current_page: number;
  total_pages: number;
  total_size: number;
}

 
export interface ApiResponseList {
  status: string;
  message: string;
  pagination: PaginationProp;
  data: UserProp[];
}

export interface PostProp {
  id: string;
  userId: string;
  title: string;
  body: string;
  createdAt: string;
}

export interface ApiResponse<T> {
  status: string;
  message: string;
  data: T;
}