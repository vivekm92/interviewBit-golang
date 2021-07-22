
bool Solution::hotel(vector<int> &arrive, vector<int> &depart, int K) {
    
    int n = arrive.size();
    vector<pair<int, int> > arr;
    for (int i = 0; i < n; i++) {
        arr.push_back(make_pair(arrive[i], 1));
    }

    for (int i = 0; i < n; i++) {
        arr.push_back(make_pair(depart[i], -1));
    }

    sort(arr.begin(), arr.end());
    int req = 0, rooms = 0;
    for (int i = 0; i < arr.size(); i++) {
        if (arr[i].second == 1) {
            rooms++;
        } else {
            rooms--;
        }

        req = max(req, rooms);
    }

    return req <= K;
}
