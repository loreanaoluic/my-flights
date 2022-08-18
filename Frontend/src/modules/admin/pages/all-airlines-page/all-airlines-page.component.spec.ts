import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllAirlinesPageComponent } from './all-airlines-page.component';

describe('AllAirlinesPageComponent', () => {
  let component: AllAirlinesPageComponent;
  let fixture: ComponentFixture<AllAirlinesPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllAirlinesPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllAirlinesPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
