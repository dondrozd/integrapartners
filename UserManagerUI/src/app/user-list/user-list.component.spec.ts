import { UserService } from './../user-service/user.service';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { User } from '../user-service/user';
import { of } from 'rxjs';

describe('UserListComponent', () => {
  let component: UserListComponent;
  let fixture: ComponentFixture<UserListComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;
  let httpClientSpy: { get: jasmine.Spy };

  let users: User[] = [
    {id: 1, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'},
    {id: 2, firstName: 'fnB', lastName: 'lnB', email: 'b@dn.com', userName: 'unB', status: 'A', department: 'depB'},
  ];


  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserListComponent ],
      providers: [{provide: UserService, useValue: userServiceSpy}]
    })
    .compileComponents();
  });

  beforeEach(() => {
    userServiceSpy = jasmine.createSpyObj(UserService, ['getUsers']);
    fixture = TestBed.createComponent(UserListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should show users values', () => {
    userServiceSpy.getUsers.and.returnValue(of(users));
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });


});
