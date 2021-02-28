import { TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';

import { UserService } from './user.service';

describe('UserService', () => {
  let service: UserService;
  let httpTestingController : HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [UserService]
    });
    service = TestBed.inject(UserService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getUsers', () => {
    it('call the url to get all users', () => {
      service.getUsers().subscribe();
      const controller = httpTestingController.expectOne("api/users");
      httpTestingController.verify();
      expect(controller.request.method).toBe('GET');
    });
  });

  describe('getUser', () => {
    it('call the url to get single user', () => {
      service.getUser(1).subscribe();
      const controller = httpTestingController.expectOne("api/users/1");
      httpTestingController.verify();
      expect(controller.request.method).toBe('GET');
    });
  });





});
