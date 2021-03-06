import { TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpEvent, HttpEventType, HttpClient } from '@angular/common/http';

import { UserService } from './user.service';
import { User } from '../models/user';

describe('UserService', () => {
  let service: UserService;
  let httpTestingController: HttpTestingController;
  let httpClient: HttpClient;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [UserService]
    });
    service = TestBed.inject(UserService);
    httpTestingController = TestBed.inject(HttpTestingController);
    httpClient = TestBed.inject(HttpClient);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getUsers', () => {
    it('call the url to get all users', () => {
      service.getUsers().subscribe();
      const controller = httpTestingController.expectOne('api/users');
      httpTestingController.verify();
      expect(controller.request.method).toBe('GET');
    });
  });

  describe('getUser', () => {
    it('call the url to get single user', () => {
      service.getUser(1).subscribe();
      const controller = httpTestingController.expectOne('api/users/1');
      httpTestingController.verify();
      expect(controller.request.method).toBe('GET');
    });
  });

  describe('deleteUser', () => {
    it('call the url to get single user', () => {
      service.deleteUser(2).subscribe();
      const controller = httpTestingController.expectOne('api/users/2');
      httpTestingController.verify();
      expect(controller.request.method).toBe('DELETE');
    });
  });

  describe('addUser', () => {
    it('call the url to add a single user', () => {
      const user: User = {id: null, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'};
      service.addUser(user).subscribe();
      const controller = httpTestingController.expectOne('api/users');
      httpTestingController.verify();
      expect(controller.request.method).toBe('POST');
    });
  });

  describe('updateUser', () => {
    it('call the url to update a single user', () => {
      const user: User = {id: 3, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'};
      service.updateUser(user).subscribe();
      const controller = httpTestingController.expectOne('api/users/3');
      httpTestingController.verify();
      expect(controller.request.method).toBe('PUT');
    });
  });

});
