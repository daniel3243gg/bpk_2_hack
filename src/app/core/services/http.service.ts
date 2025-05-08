import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';
import { ApiResponse } from '../models/api-response.model';

@Injectable({
  providedIn: 'root'
})
export class HttpService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  getData<T>(endpoint: string, params?: any): Observable<ApiResponse<T>> {
    const httpParams = params ? new HttpParams({ fromObject: params }) : undefined;
    return this.http.get<ApiResponse<T>>(`${this.baseUrl}${endpoint}`, { params: httpParams });
  }

  postData<T>(endpoint: string, data?: any, params?: any): Observable<ApiResponse<T>> {
    const httpParams = params ? new HttpParams({ fromObject: params }) : undefined;
    return this.http.post<ApiResponse<T>>(`${this.baseUrl}${endpoint}`, data, { params: httpParams });
  }

  updateData<T>(endpoint: string, id?: number, data?: any, params?: any): Observable<ApiResponse<T>> {
    const url = id ? `${this.baseUrl}${endpoint}/${id}` : `${this.baseUrl}/${endpoint}`;
    const httpParams = params ? new HttpParams({ fromObject: params }) : undefined;
    return this.http.put<ApiResponse<T>>(url, data, { params: httpParams });
  }

  deleteData<T>(endpoint: string, id?: number, params?: any): Observable<ApiResponse<T>> {
    const url = id ? `${this.baseUrl}${endpoint}/${id}` : `${this.baseUrl}/${endpoint}`;
    const httpParams = params ? new HttpParams({ fromObject: params }) : undefined;
    return this.http.delete<ApiResponse<T>>(url, { params: httpParams });
  }
}
