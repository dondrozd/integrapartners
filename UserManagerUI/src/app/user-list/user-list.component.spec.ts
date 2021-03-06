import { UserService } from './../user-service/user.service';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { User } from '../models/user';
import { Observable, of, throwError } from 'rxjs';

describe('UserListComponent', () => {
  let component: UserListComponent;
  let fixture: ComponentFixture<UserListComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;

  const users: User[] = [
    {id: 1, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'},
    {id: 2, firstName: 'fnB', lastName: 'lnB', email: 'b@dn.com', userName: 'unB', status: 'A', department: 'depB'},
  ];


  beforeEach(async () => {
    userServiceSpy = jasmine.createSpyObj(UserService, ['getUsers']);
    await TestBed.configureTestingModule({
      declarations: [ UserListComponent ],
      providers: [ { provide: UserService, useValue: userServiceSpy } ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserListComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    userServiceSpy.getUsers.and.returnValue(of(users));
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should show users values', () => {
    userServiceSpy.getUsers.and.returnValue(of(users));
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should set hasError true on http error ', () => {
    userServiceSpy.getUsers.and.returnValue(throwError({status: 500}));
    fixture.detectChanges();
    expect(component.hasError).toBe(true);
  });


});
