import { MatSnackBar } from '@angular/material/snack-bar';
import { RouterTestingModule } from '@angular/router/testing';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';
import { UserService } from './../user-service/user.service';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserComponent } from './user.component';
import { User } from './../models/user';
import { of } from 'rxjs';

describe('UserComponent', () => {

  let component: UserComponent;
  let fixture: ComponentFixture<UserComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;
  let routerSpy: jasmine.SpyObj<Router>;
  let paramMapSpy: jasmine.SpyObj<ParamMap>;
  let snackBarSpy: jasmine.SpyObj<MatSnackBar>;

  const user: User = {id: 1, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'};

  beforeEach(async () => {
    userServiceSpy = jasmine.createSpyObj(UserService, ['getUser', 'addUser', 'updateUser']);
    routerSpy = jasmine.createSpyObj('Router', ['navigate']);
    paramMapSpy = jasmine.createSpyObj('ParamMap', ['get', 'has']);
    snackBarSpy = jasmine.createSpyObj('MatSnackBar', ['open']);
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule ],
      declarations: [ UserComponent ],
      providers: [
        { provide: UserService, useValue: userServiceSpy },
        { provide: ActivatedRoute, useValue: {snapshot: { paramMap: paramMapSpy }} },
        { provide: Router, useValue: routerSpy },
        { provide: MatSnackBar, useValue: snackBarSpy }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserComponent);
    component = fixture.componentInstance;
    paramMapSpy.has.and.returnValue(true);
    paramMapSpy.get.and.returnValue('1');
  });

  it('should create', () => {
    userServiceSpy.getUser.and.returnValue(of(user));
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should route on cancel', () => {
    userServiceSpy.getUser.and.returnValue(of(user));
    fixture.detectChanges();
    component.onClickCancel();
    const [actualPath] = routerSpy.navigate.calls.first().args;
    expect(actualPath).toEqual(['users']);
  });

  it('should update existing on save', () => {
    userServiceSpy.getUser.and.returnValue(of(user));
    userServiceSpy.updateUser.and.returnValue(of());
    fixture.detectChanges();
    component.onClickSave();
    expect(userServiceSpy.updateUser).toHaveBeenCalled();
  });

  it('should add new user on save', () => {
    paramMapSpy.has.and.returnValue(false);
    userServiceSpy.addUser.and.returnValue(of());
    fixture.detectChanges();
    component.onClickSave();
    expect(userServiceSpy.addUser).toHaveBeenCalled();
  });
});
