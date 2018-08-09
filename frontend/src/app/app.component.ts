import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {models} from "./shared/types/proto-types";
import User = models.User;
import {map} from "rxjs/operators";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    title = 'ng6-app';

    constructor(private http: HttpClient) {
        this.http.get("api/user")
            .pipe(
                map((res: any) => JSON.parse(res))
            )
            .subscribe((users: User[]) => {
                console.log(users);
            });

        let user: User = new User();
        user.Name = "Daniel";
        user.Phone = "1234";
        this.http.post("api/user", user)
            .pipe(
                map((res: any) => JSON.parse(res))
            )
            .subscribe((user: User) => {
                console.log(user);
            });
    }
}
